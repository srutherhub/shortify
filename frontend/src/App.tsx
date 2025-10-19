import { Route, Routes } from "react-router-dom";
import "./App.css";
import LoginPage from "./routes/LoginPage";
import HankoProfile from "./components/auth/HankoProfile";
import HankoPrivateRoute from "./components/auth/HankoPrivateRoute";
import Navbar from "./components/nav/Navbar";

function App() {
  return (
    <div>
      <Navbar />
      <Routes>
        <Route path="/" element={<h1>Hello!</h1>}></Route>
        <Route path="/login" element={<LoginPage />}></Route>
        <Route path="/products" element={<h1>Hello from products</h1>}></Route>
        <Route path="/pricing" element={<h1>Hello from pricing</h1>}></Route>
        <Route path="/docs" element={<h1>Hello from docs</h1>}></Route>
        <Route
          path="/dashboard"
          element={
            <HankoPrivateRoute>
              {" "}
              <h1>Hello from dashboard</h1>
            </HankoPrivateRoute>
          }
        >
          <Route path="/dashboard/profile" element={<HankoProfile />}></Route>
        </Route>
        <Route
          path="/changelog"
          element={<h1>Hello from changelog</h1>}
        ></Route>
      </Routes>
    </div>
  );
}

export default App;
