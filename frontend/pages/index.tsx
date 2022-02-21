import React, { useReducer, useEffect } from 'react';
import { Flex, Container, Heading, Stack, Image, GridItem, Grid, Skeleton, Text } from '@chakra-ui/react';

import { fetchRestaurants } from './api/restaurants';
import MainLogo from '../public/images/logo.jpg';
import MainCoverImage from '../public/images/main-cover-image.jpg';
import RestaurantImage from '../public/images/restaurant-image.jpg';
import Link from 'next/link'


import {
  initialState,
  restaurantsActionTypes,
  restaurantsReducer,
} from '../state/reducers/restaurants';
import { REQUEST_STATE } from '../state/constants';


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
      <Stack >
        {
          state.fetchState === REQUEST_STATE.LOADING ?
            <Grid templateColumns='repeat(3, 1fr)' gap={6}>
              <Skeleton w='100%' h='200' />
              <Skeleton w='100%' h='200' />
              <Skeleton w='100%' h='200' />
            </Grid>
        :
          <Grid templateColumns='repeat(3, 1fr)' gap={6}>
              {state.restaurantsList.map((restaurant, index) => (
              <GridItem w='100%' h='200' key={restaurant.id}>
                <Image src={RestaurantImage.src} alt='restaurant' boxSize='100%' />
                  <Link href={`/food/${restaurant.id}`} key={index}>
                    <a>{restaurant.name}</a>
                  </Link>
                  <Text fontSize='sm'>{`配送料：${restaurant.fee}円 ${restaurant.time_required}分`}</Text>
              </ GridItem >
              ))}
          </Grid>
        }
      </Stack>
    </Container>
  );
};
export default Index;
