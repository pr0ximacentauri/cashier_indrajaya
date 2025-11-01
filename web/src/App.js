import AppRouter from "./routes/AppRouter";

function App() {
  return (
    <Router>
      <nav
        style={{
          background: "#333",
          color: "#fff",
          padding: "1rem",
          display: "flex",
          gap: "1rem",
        }}
      >
        <Link style={{ color: "white" }} to="/">
          🛍️ Produk
        </Link>
        <Link style={{ color: "white" }} to="/kasir">
          💳 Kasir
        </Link>
      </nav>

      <Routes>
        <Route path="/" element={<ProductPage />} />
        <Route path="/kasir" element={<TransactionPage />} />
      </Routes>
    </Router>
  );
};

export default App;