import React, { useEffect, useReducer, useState } from 'react';
import {
  Container,
  Stack,
  Flex,
  Heading,
  Image,
  Skeleton,
  Grid,
  GridItem,
  Text,
  ButtonGroup,
  IconButton,
  Button,
  Modal,
  ModalOverlay,
  ModalContent,
  ModalHeader,
  ModalCloseButton,
  ModalBody,
  ModalFooter,
  useColorModeValue,
} from '@chakra-ui/react';
import { MinusIcon, AddIcon } from '@chakra-ui/icons';

import { useRouter } from 'next/router';
import { fetchFoods } from '../api/foods';
import { REQUEST_STATE } from '../../state/constants';
import { HTTP_STATUS_CODE } from '../../state/constants';
import { postLineFoods, replaceLineFoods } from '../api/line_foods';

import MainLogo from '../../public/images/logo.jpg';
import OrderHeaderImage from '../../public/images/order-header.jpg';
import FoodImage from '../../public/images/food-image.jpg';
import { Food } from '../../lib/food/type';
import { NewOrderConfirmModal } from '../../components/NewOrderConfirmModal';
import {
  initialState as foodsInitialState,
  foodsActionTypes,
  foodsReducer,
} from '../../state/reducers/foods';

export let initFood = {
  id: '',
  restaurant_id: '',
  name: '',
  price: 0,
  description: '',
};

type Props = {
  isOpenOrderDialog: boolean;
  selectedFood: Food;
  selectedFoodCount: number;
  isOpenNewOrderDialog: boolean;
  existingResutaurautName: string;
  newResutaurautName: string;
};

export const Foods = () => {
  const router = useRouter();
  const { restaurantId } = router.query;
  const [foodsState, dispatch] = useReducer(foodsReducer, foodsInitialState);
  const initialState: Props = {
    isOpenOrderDialog: false,
    selectedFood: initFood,
    selectedFoodCount: 1,
    isOpenNewOrderDialog: false,
    existingResutaurautName: '',
    newResutaurautName: '',
  };
  const [state, setState] = useState(initialState);

  const onClickFood = (food: Food) => {
    setState({
      ...state,
      isOpenOrderDialog: true,
      selectedFood: food,
    });
  };

  const onClickCountUp = (state: Props) => {
    setState({
      ...state,
      selectedFoodCount: state.selectedFoodCount + 1,
    });
  };
  const onClickCountDown = (state: Props) => {
    setState({
      ...state,
      selectedFoodCount: state.selectedFoodCount - 1,
    });
  };

  useEffect(() => {
    fetchFoods(restaurantId as string).then((data) => {
      dispatch({
        type: foodsActionTypes.FETCH_SUCCESS,
        payload: {
          foods: data.foods,
        },
      });
    });
  }, []);

  const submitOrder = () => {
    postLineFoods({
      foodId: state.selectedFood.id,
      restaurantId: state.selectedFood.restaurant_id,
      count: state.selectedFoodCount,
    })
      .then(() => router.push('/orders'))
      .catch((e) => {
        if (e.response.status === HTTP_STATUS_CODE.NOT_ACCEPTABLE) {
          setState({
            ...state,
            isOpenOrderDialog: false,
            isOpenNewOrderDialog: true,
            existingResutaurautName: e.response.data.existing_restaurant,
            newResutaurautName: e.response.data.new_restaurant,
          });
        } else {
          throw e;
        }
      });
  };

  const replaceOrder = () => {
    replaceLineFoods({
      foodId: state.selectedFood.id,
      restaurantId: state.selectedFood.restaurant_id,
      count: state.selectedFoodCount,
    }).then(() => router.push('/orders'));
  };

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
      <Stack>
        {foodsState.fetchState === REQUEST_STATE.LOADING ? (
          <Grid templateColumns='repeat(4, 1fr)' gap={6}>
            <Skeleton w='100%' h='200' />
            <Skeleton w='100%' h='200' />
            <Skeleton w='100%' h='200' />
            <Skeleton w='100%' h='200' />
            <Skeleton w='100%' h='200' />
            <Skeleton w='100%' h='200' />
            <Skeleton w='100%' h='200' />
            <Skeleton w='100%' h='200' />
            <Skeleton w='100%' h='200' />
            <Skeleton w='100%' h='200' />
            <Skeleton w='100%' h='200' />
            <Skeleton w='100%' h='200' />
          </Grid>
        ) : (
          <Grid templateColumns='repeat(4, 1fr)' gap={10}>
            {foodsState.foodsList.map((food, index) => (
              <GridItem w='100%' h='25%' key={food.id}>
                <Stack
                  borderWidth='1px'
                  borderRadius='lg'
                  w={{ sm: '100%', md: '100%' }}
                  height={{ sm: '476px', md: '15rem' }}
                  direction={{ base: 'column', md: 'row' }}
                  bg={useColorModeValue('white', 'gray.900')}
                  boxShadow={'2xl'}
                  padding={4}
                  key={food.id}
                >
                  <Flex flex={1} bg='black'>
                    <Image
                      key={food.id}
                      objectFit='contain'
                      boxSize='200'
                      src={FoodImage.src}
                    />
                  </Flex>
                  <Stack
                    flex={1}
                    flexDirection='column'
                    alignItems='center'
                    p={1}
                    pt={2}
                  >
                    <Heading
                      fontSize={'1xl'}
                      fontFamily={'body'}
                      as={'button'}
                      _hover={{
                        background: 'white',
                        color: 'teal.500',
                      }}
                      onClick={() => onClickFood(food)}
                    >
                      {food.name}
                    </Heading>
                    <Text
                      textAlign={'center'}
                      color={useColorModeValue('gray.700', 'gray.400')}
                      fontSize='sm'
                    >
                      {food.description}
                    </Text>
                    <Text
                      textAlign={'left'}
                      color={useColorModeValue('gray.700', 'gray.400')}
                      fontSize='md'
                    >
                      ¥{food.price}
                    </Text>
                  </Stack>
                </Stack>
              </GridItem>
            ))}
          </Grid>
        )}
      </Stack>
      {state.isOpenOrderDialog && (
        <Modal
          isOpen={state.isOpenOrderDialog}
          onClose={() =>
            setState({
              ...state,
              isOpenOrderDialog: false,
            })
          }
        >
          <ModalOverlay />
          <ModalContent>
            <ModalHeader>購入する商品数を入力してください。</ModalHeader>
            <ModalCloseButton />
            <ModalBody>
              <Image src={OrderHeaderImage.src} alt='order' maxW='100%' />
              <Text
                textAlign={'center'}
                color={useColorModeValue('gray.700', 'gray.400')}
                fontSize='sm'
              >
                {state.selectedFood.name}
              </Text>
              <Text
                textAlign={'center'}
                color={useColorModeValue('gray.700', 'gray.400')}
                fontSize='sm'
              >
                {state.selectedFood.description}
              </Text>
            </ModalBody>
            <ModalFooter>
              <Text
                textAlign={'center'}
                color={useColorModeValue('gray.700', 'gray.400')}
                fontSize='sm'
              >
                合計{state.selectedFoodCount}個
              </Text>
              <ButtonGroup size='sm' isAttached variant='outline'>
                <IconButton
                  aria-label='Add to friends'
                  isDisabled={state.selectedFoodCount <= 1}
                  icon={<MinusIcon onClick={() => onClickCountDown(state)} />}
                />
                <IconButton
                  aria-label='Add to friends'
                  isDisabled={state.selectedFoodCount >= 9}
                  icon={<AddIcon onClick={() => onClickCountUp(state)} />}
                />
              </ButtonGroup>
              <Stack direction='row' align='center'>
                <Button
                  size='md'
                  height='40px'
                  width='200px'
                  border='2px'
                  onClick={() => submitOrder()}
                >
                  {`${state.selectedFoodCount}点を注文に追加`}{' '}
                  {`¥${state.selectedFoodCount * state.selectedFood.price}`}
                </Button>
              </Stack>
            </ModalFooter>
          </ModalContent>
        </Modal>
      )}
      {state.isOpenNewOrderDialog && (
        <NewOrderConfirmModal
          isOpen={state.isOpenNewOrderDialog}
          onClose={() => setState({ ...state, isOpenNewOrderDialog: false })}
          existingResutaurautName={state.existingResutaurautName}
          newResutaurautName={state.newResutaurautName}
          onClickSubmit={() => replaceOrder()}
        />
      )}
    </Container>
  );
};

export default Foods;
