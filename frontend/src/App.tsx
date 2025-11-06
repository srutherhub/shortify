import { Route, Routes } from "react-router-dom";
import "./App.css";
import LoginPage from "./routes/LoginPage";
import HankoProfile from "./components/auth/HankoProfile";
import HankoPrivateRoute from "./components/auth/HankoPrivateRoute";
import Navbar from "./components/nav/Navbar";
import AppPage from "./routes/AppPage";
import DashboardPage from "./routes/DashboardPage";

function App() {
  return (
    <div>
      <Navbar />
      <Routes>
        <Route path="/" element={<h1>Hello!</h1>}></Route>
        <Route path="/login" element={<LoginPage />}></Route>
        <Route path="/products" element={<h1>Hello from products</h1>}></Route>
        <Route path="/pricing" element={<h1>Hello from pricing</h1>}></Route>
        <Route path="/docs" element={<h1>Documents</h1>}></Route>
        <Route
          path="/app"
          element={
            <HankoPrivateRoute>
              <AppPage />
            </HankoPrivateRoute>
          }
        >
          <Route index element={<DashboardPage />} />
          <Route path="profile" element={<HankoProfile />}></Route>
          <Route path="analytics" element={<h1>Analytics</h1>}></Route>
          <Route path="manage" element={<h1>Manage Links</h1>}></Route>
        </Route>
        <Route path="/changelog" element={<h1>Changelog</h1>}></Route>
      </Routes>
    </div>
  );
}

export default App;
