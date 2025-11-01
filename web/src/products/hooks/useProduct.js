import { useState, useEffect } from "react";
import * as productService from "../services/productService";

export const useProducts = () => {
  const [products, setProducts] = useState([]);
  const [loading, setLoading] = useState(true);

  const fetchProducts = async () => {
    setLoading(true);
    const data = await productService.getProducts();
    setProducts(data);
    setLoading(false);
  };

  const addProduct = async (product) => {
    await productService.addProduct(product);
    fetchProducts();
  };

  const updateProduct = async (id, product) => {
    await productService.updateProduct(id, product);
    fetchProducts();
  };

  const deleteProduct = async (id) => {
    await productService.deleteProduct(id);
    fetchProducts();
  };

  useEffect(() => {
    fetchProducts();
  }, []);

  return {
    products,
    loading,
    addProduct,
    updateProduct,
    deleteProduct,
  };
};
