import { REQUEST_STATE } from '../constants';
import { LineFoodResponse } from '../../lib/line_foods/type';

export const initRestaurant = {
  id: '',
  name: '',
  fee: 0,
  time_required: 0,
};

export const initialLineFood: LineFoodResponse = {
  line_food_ids: [],
  restaurant: initRestaurant,
  count: 0,
  amount: 0,
};

export const initialState: LineFoodsState = {
  lineFoodsSummary: initialLineFood,
  fetchState: REQUEST_STATE.INITIAL,
  postState: REQUEST_STATE.INITIAL,
};

export const lineFoodsActionTypes = {
  FETCHING: 'FETCHING',
  FETCH_SUCCESS: 'FETCH_SUCCESS',
  POSTING: 'POSTING',
  POST_SUCCESS: 'POST_SUCCESS',
};

type LineFoodsState =
  | {
      lineFoodsSummary: LineFoodResponse;
      fetchState: string;
      postState?: string;
    }
  | {
      lineFoodsSummary: LineFoodResponse;
      postState: string;
      fetchState?: string;
    };

type ActionType = {
  type: string;
  payload: {
    lineFoodsSummary: LineFoodResponse;
  };
};

export const lineFoodsReducer = (state: LineFoodsState, action: ActionType) => {
  switch (action.type) {
    case lineFoodsActionTypes.FETCHING:
      return {
        ...state,
        fetchState: REQUEST_STATE.LOADING,
      };
    case lineFoodsActionTypes.FETCH_SUCCESS:
      return {
        fetchState: REQUEST_STATE.OK,
        lineFoodsSummary: action.payload.lineFoodsSummary,
      };
    case lineFoodsActionTypes.POSTING:
      return {
        ...state,
        postState: REQUEST_STATE.LOADING,
      };
    case lineFoodsActionTypes.POST_SUCCESS:
      return {
        ...state,
        postState: REQUEST_STATE.OK,
      };
    default:
      throw new Error();
  }
};
// --- ここまで追加 ---
