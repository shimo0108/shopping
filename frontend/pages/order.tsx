import React, { useEffect, useReducer } from 'react';
import {
  initialState,
  lineFoodsActionTypes,
  lineFoodsReducer,
  initialLineFood,
} from '../state/reducers/lineFoods';
import {
  Stack,
  Box,
  Flex,
  Heading,
  Text,
  useColorModeValue,
  Grid,
  GridItem,
} from '@chakra-ui/react';
import Link from 'next/link';
import { MdWork, MdAccessTimeFilled } from 'react-icons/md';
import { postOrder } from './api/order';
import { fetchLineFoods } from './api/line_foods';
import { REQUEST_STATE } from '../state/constants';
import { SmallCloseIcon } from '@chakra-ui/icons';

export const Order = () => {
  const [state, dispatch] = useReducer(lineFoodsReducer, initialState);

  useEffect(() => {
    dispatch({
      type: lineFoodsActionTypes.FETCHING,
      payload: {
        lineFoodsSummary: initialLineFood,
      },
    });
    fetchLineFoods()
      .then((data) =>
        dispatch({
          type: lineFoodsActionTypes.FETCH_SUCCESS,
          payload: {
            lineFoodsSummary: data,
          },
        })
      )
      .catch((e) => console.error(e));
  }, []);

  const postLineFoods = () => {
    dispatch({
      type: lineFoodsActionTypes.POSTING,
      payload: {
        lineFoodsSummary: state.lineFoodsSummary,
      },
    });
    postOrder({
      line_food_ids: state.lineFoodsSummary.line_food_ids,
    }).then(() => {
      dispatch({
        type: lineFoodsActionTypes.POST_SUCCESS,
        payload: {
          lineFoodsSummary: state.lineFoodsSummary,
        },
      });
      window.location.reload();
    });
  };

  const orderButtonLabel = () => {
    switch (state.postState) {
      case REQUEST_STATE.LOADING:
        return '注文中...';
      case REQUEST_STATE.OK:
        return '注文が完了しました！';
      default:
        return '注文を確定する';
    }
  };

  return (
    <Flex
      minH={'100vh'}
      align={'center'}
      justify={'center'}
      bg={useColorModeValue('gray.50', 'gray.800')}
    >
      <Stack
        spacing={4}
        w={'full'}
        maxW={'md'}
        bg={useColorModeValue('white', 'gray.700')}
        rounded={'xl'}
        boxShadow={'lg'}
        p={6}
        my={12}
      >
        <Heading lineHeight={1.1} fontSize={{ base: '2xl', sm: '3xl' }}>
          注文確認
        </Heading>
        <Grid templateColumns='repeat(5, 1fr)' gap={4}>
          <GridItem>
            <MdWork size={30}></MdWork>
          </GridItem>
          <GridItem colStart={4} colEnd={6} h='10'>
            <Link href={`/foods/${state.lineFoodsSummary.restaurant.id}`}>
              <a>{state.lineFoodsSummary.restaurant.name}</a>
            </Link>
          </GridItem>
        </Grid>
        <Grid templateColumns='repeat(5, 1fr)' gap={4}>
          <GridItem>
            <MdAccessTimeFilled size={30}></MdAccessTimeFilled>
          </GridItem>
          <GridItem colStart={4} colEnd={6} h='10'>
            <Text fontSize='md'>
              {state.lineFoodsSummary.restaurant.time_required}分で到着予定
            </Text>
          </GridItem>
        </Grid>
        <Grid templateColumns='repeat(5, 1fr)' gap={4}>
          <GridItem>
            <Text fontSize='md'>商品数</Text>
          </GridItem>
          <GridItem colStart={4} colEnd={6} h='10'>
            <Text fontSize='md'>{state.lineFoodsSummary.count}</Text>
          </GridItem>
        </Grid>
        <Grid templateColumns='repeat(5, 1fr)' gap={4}>
          <GridItem colStart={1} colEnd={3} h='10'>
            <Text fontSize='md'>商品数:{state.lineFoodsSummary.count}</Text>
          </GridItem>
          <GridItem colStart={4} colEnd={6} h='10'>
            <Text fontSize='md'>¥{state.lineFoodsSummary.amount}</Text>
          </GridItem>
        </Grid>
        <Grid templateColumns='repeat(5, 1fr)' gap={4}>
          <GridItem colStart={1} colEnd={3} h='10'>
            <Text fontSize='md'>配送料</Text>
          </GridItem>
          <GridItem colStart={4} colEnd={6} h='10'>
            <Text fontSize='md'>¥{state.lineFoodsSummary.restaurant.fee}</Text>
          </GridItem>
        </Grid>
        <Grid templateColumns='repeat(5, 1fr)' gap={4}>
          <GridItem colStart={1} colEnd={3} h='10'>
            <Text fontSize='md'>合計</Text>
          </GridItem>
          <GridItem colStart={4} colEnd={6} h='10'>
            <Text fontSize='md'>
              ¥
              {state.lineFoodsSummary.amount +
                state.lineFoodsSummary.restaurant.fee}
            </Text>
          </GridItem>
        </Grid>
        <Box
          as='button'
          height='40px'
          lineHeight='1.2'
          transition='all 0.2s cubic-bezier(.08,.52,.52,1)'
          border='1px'
          px='8px'
          borderRadius='2px'
          fontSize='14px'
          fontWeight='semibold'
          bg='black'
          borderColor='#ccd0d5'
          color='white'
          _hover={{ bg: '#ebedf0' }}
          _active={{
            bg: '#dddfe2',
            transform: 'scale(0.98)',
            borderColor: '#bec3c9',
          }}
          _focus={{
            boxShadow:
              '0 0 1px 2px rgba(88, 144, 255, .75), 0 1px 1px rgba(0, 0, 0, .15)',
          }}
          onClick={() => postLineFoods()}
        >
          注文を確定する
        </Box>
      </Stack>
    </Flex>
  );
};

export default Order;
