import { BrowserRouter as Router, Routes,Route } from "react-router-dom";
import { Login } from "./Components/Login";
// import { Navbar } from "./Components/Navbar";
import { DashCreateGRN } from "./Pages/DashCreateGRN";
function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<Login/>}/>
        <Route path="/createGRN" element={<DashCreateGRN/>}/>
      </Routes>
    </Router>
  );
}

export default App;
