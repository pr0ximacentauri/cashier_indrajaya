import { useState } from "react";
import { useProducts } from "../hooks/useProducts";
import { ProductForm } from "../components/ProductForm";
import { ProductTable } from "../components/ProductTable";

const ProductPage = () => {
  const { products, addProduct, updateProduct, deleteProduct, loading } = useProducts();
  const [editingProduct, setEditingProduct] = useState(null);

  const handleSubmit = async (form, id) => {
    if (id) await updateProduct(id, form);
    else await addProduct(form);
    setEditingProduct(null);
  };

  if (loading) return <p>Loading...</p>;

  return (
    <div style={{ padding: "2rem" }}>
      <h1>🛒 Manajemen Produk</h1>
      <ProductForm onSubmit={handleSubmit} editingProduct={editingProduct} />
      <ProductTable products={products} onEdit={setEditingProduct} onDelete={deleteProduct} />
    </div>
  );
};

export default ProductPage;
