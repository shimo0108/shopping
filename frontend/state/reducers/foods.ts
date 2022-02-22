import { REQUEST_STATE } from '../constants';
import { Food } from '../../lib/food/type';

export const initialState = {
  fetchState: REQUEST_STATE.INITIAL,
  foodsList: [],
};

export const foodsActionTypes = {
  FETCHING: 'FETCHING',
  FETCH_SUCCESS: 'FETCH_SUCCESS',
};

type FoodState = {
  fetchState: string;
  foodsList: Food[];
};

type ActionType = {
  type: string;
  payload: {
    foods: [];
  };
};

export const foodsReducer = (state: FoodState, action: ActionType) => {
  switch (action.type) {
    case foodsActionTypes.FETCHING:
      return {
        ...state,
        fetchState: REQUEST_STATE.LOADING,
      };
    case foodsActionTypes.FETCH_SUCCESS:
      return {
        fetchState: REQUEST_STATE.OK,
        foodsList: action.payload.foods,
      };
    default:
      throw new Error();
  }
};
