import Restaurant from '../restaurant/type';

export type LineFood = {
  id: string;
  restaurant_id: string;
  food_id: string;
  order_id: string;
  count: number;
  active: boolean;
};

export type LineFoodResponse = {
  line_food_ids: string[];
  restaurant: Restaurant;
  count: number;
  amount: number;
};
