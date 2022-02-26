// --- ここから追加 ---
import axios from 'axios';
import { lineFoodCreate, lineFoodsReplace, lineFoods } from './urls/index';

type PostProps = {
  restaurantId: string;
  foodId: string;
  count: number;
};

export const postLineFoods = (params: PostProps) => {
  return axios
    .post(lineFoodCreate, {
      restaurant_id: params.restaurantId,
      food_id: params.foodId,
      count: params.count,
    })
    .then((res) => {
      return res.data;
    })
    .catch((e) => {
      throw e;
    });
};

type ReplaceProps = {
  restaurantId: string;
  foodId: string;
  count: number;
};
export const replaceLineFoods = (params: ReplaceProps) => {
  return axios
    .put(lineFoodsReplace(params.restaurantId), {
      food_id: params.foodId,
      restaurantId: params.restaurantId,
      count: params.count,
    })
    .then((res) => {
      return res.data;
    })
    .catch((e) => {
      throw e;
    });
};

export const fetchLineFoods = () => {
  return axios
    .get(lineFoods)
    .then((res) => {
      return res.data;
    })
    .catch((e) => {
      throw e;
    });
};
