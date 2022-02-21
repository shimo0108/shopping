const DEFAULT_API_LOCALHOST = 'http://localhost:9999/api/v1';

export const restaurantsIndex = `${DEFAULT_API_LOCALHOST}/restaurants`;
export const foodsIndex = (restaurantId: string) =>
  `${DEFAULT_API_LOCALHOST}/foods/${restaurantId}`;
export const lineFoods = `${DEFAULT_API_LOCALHOST}/line_foods`;
export const lineFoodsReplace = (restaurantId: string) =>
  `${DEFAULT_API_LOCALHOST}/line_foods/${restaurantId}`;
export const orders = `${DEFAULT_API_LOCALHOST}/orders`;
