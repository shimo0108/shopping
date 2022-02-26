import axios from 'axios';
import { orders } from './urls/index';
type Props = {
  line_food_ids: string[];
};

export const postOrder = (params: Props) => {
  return axios
    .post(orders, {
      line_food_ids: params.line_food_ids,
    })
    .then((res) => {
      return res.data;
    })
    .catch((e) => console.error(e));
};
