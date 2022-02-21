import { REQUEST_STATE } from '../constants';
import { Restaurant } from '../../lib/restaurant/type';

export const initialState: RestaurantState = {
  fetchState: REQUEST_STATE.INITIAL,
  restaurantsList: [],
};

export const restaurantsActionTypes = {
  FETCHING: 'FETCHING',
  FETCH_SUCCESS: 'FETCH_SUCCESS',
};

type RestaurantState = {
  fetchState: string;
  restaurantsList: Restaurant[];
};

type ActionType = {
  type: string;
  payload: {
    restaurants: [];
  };
};

export const restaurantsReducer = (
  state: RestaurantState,
  action: ActionType
) => {
  switch (action.type) {
    case restaurantsActionTypes.FETCHING:
      return {
        ...state,
        fetchState: REQUEST_STATE.LOADING,
      };
    case restaurantsActionTypes.FETCH_SUCCESS:
      return {
        fetchState: REQUEST_STATE.OK,
        restaurantsList: action.payload.restaurants,
      };
    default:
      throw new Error();
  }
};
