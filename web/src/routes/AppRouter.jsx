import { BrowserRouter as Router, Routes, Route, Navigate } from "react-router-dom";
import ProductPage from "../features/products/pages/ProductPage";

const AppRouter = () => (
  <Router>
    <Routes>
      <Route path="/" element={<Navigate to="/products" />} />
      <Route path="/products" element={<ProductPage />} />
    </Routes>
  </Router>
);

export default AppRouter;
