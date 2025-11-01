import { useState, useEffect } from "react";

export const ProductForm = ({ onSubmit, editingProduct }) => {
  const [form, setForm] = useState({
    name: "",
    category: "",
    price: "",
    stock: "",
  });

  useEffect(() => {
    if (editingProduct) setForm(editingProduct);
  }, [editingProduct]);

  const handleChange = (e) => {
    setForm({ ...form, [e.target.name]: e.target.value });
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    onSubmit(form, editingProduct?.id);
    setForm({ name: "", category: "", price: "", stock: "" });
  };

  return (
    <form onSubmit={handleSubmit} style={{ marginBottom: "1.5rem" }}>
      <input name="name" value={form.name} onChange={handleChange} placeholder="Nama Produk" required />{" "}
      <input name="category" value={form.category} onChange={handleChange} placeholder="Kategori" required />{" "}
      <input name="price" type="number" value={form.price} onChange={handleChange} placeholder="Harga" required />{" "}
      <input name="stock" type="number" value={form.stock} onChange={handleChange} placeholder="Stok" required />{" "}
      <button type="submit">{editingProduct ? "💾 Simpan" : "➕ Tambah"}</button>
    </form>
  );
};
