const ProductItem = ({ product, onEdit, onDelete }) => {
  return (
    <tr>
      <td>{product.name}</td>
      <td>{product.category}</td>
      <td>Rp {Number(product.price).toLocaleString()}</td>
      <td>{product.stock}</td>
      <td>
        <button onClick={() => onEdit(product)}>✏️ Edit</button>{" "}
        <button onClick={() => onDelete(product.id)}>🗑️ Hapus</button>
      </td>
    </tr>
  );
};

export default ProductItem;
