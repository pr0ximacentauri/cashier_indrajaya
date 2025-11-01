import axios from "axios";

const API_URL = "/api/products";

export const getProducts = async () => {
  const res = await axios.get(API_URL);
  return res.data;
};

export const addProduct = async (data) => {
  const res = await axios.post(API_URL, data);
  return res.data;
};

export const updateProduct = async (id, data) => {
  const res = await axios.put(`${API_URL}/${id}`, data);
  return res.data;
};

export const deleteProduct = async (id) => {
  await axios.delete(`${API_URL}/${id}`);
};
