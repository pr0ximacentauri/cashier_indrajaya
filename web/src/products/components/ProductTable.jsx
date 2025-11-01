import ProductItem from "./ProductItem";

export const ProductTable = ({ products, onEdit, onDelete }) => {
  return (
    <table border="1" cellPadding="10" width="100%">
      <thead>
        <tr>
          <th>Nama</th>
          <th>Kategori</th>
          <th>Harga</th>
          <th>Stok</th>
          <th>Aksi</th>
        </tr>
      </thead>
      <tbody>
        {products.map((p) => (
          <ProductItem
            key={p.id}
            product={p}
            onEdit={onEdit}
            onDelete={onDelete}
          />
        ))}
      </tbody>
    </table>
  );
};
