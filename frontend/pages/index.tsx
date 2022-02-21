import React, { useReducer, useEffect } from 'react';
import { Flex, Container, Heading, Stack, Image } from '@chakra-ui/react';

import { fetchRestaurants } from './api/restaurants';
import MainLogo from '../public/images/logo.jpg';
import MainCoverImage from '../public/images/main-cover-image.jpg';
import {
  initialState,
  restaurantsActionTypes,
  restaurantsReducer,
} from '../state/reducers/restaurants';

export const Index = () => {
  const [state, dispatch] = useReducer(restaurantsReducer, initialState);

  useEffect(() => {
    dispatch({
      type: restaurantsActionTypes.FETCHING,
      payload: {
        restaurants: [],
      },
    });
    fetchRestaurants().then((data) =>
      dispatch({
        type: restaurantsActionTypes.FETCH_SUCCESS,
        payload: {
          restaurants: data.restaurants,
        },
      })
    );
  }, []);
  return (
    <Container maxW='100%'>
      <Flex
        as='nav'
        align='center'
        justify='space-between'
        wrap='wrap'
        padding={3}
        bg='white'
        color='white'
      >
        <Flex align='center' mr={5}>
          <Heading as='h1' size='md' letterSpacing={'tighter'}>
            <Image src={MainLogo.src} alt='logo' h='70px' />
          </Heading>
        </Flex>
      </Flex>
      <Stack textAlign={'center'} align={'center'}>
        <Image src={MainCoverImage.src} alt='main' boxSize='100%' />
      </Stack>
      {state.restaurantsList.map((restaurant) => (
        <div key={restaurant.id}>{restaurant.name}</div>
      ))}
    </Container>
  );
};
export default Index;
