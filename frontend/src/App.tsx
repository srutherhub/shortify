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
        <Route
          path="/profile"
          element={
            <HankoPrivateRoute>
              <HankoProfile />
            </HankoPrivateRoute>
          }
        ></Route>
      </Routes>
    </div>
  );
}

export default App;
