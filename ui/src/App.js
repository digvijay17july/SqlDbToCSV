import logo from './logo.svg';
import './App.css';
import Welcome from './welcome/welcome';
import {
  BrowserRouter as Router,
  Route,
  Routes
} from "react-router-dom";
import Migration from './migration/migration';
function App() {
  return (

  <Router>
  <Routes>
    <Route exact path='/' element={<Welcome > </Welcome>}> </Route>
  </Routes>
  <Routes>
    <Route exact path='/migration' element={<Migration > </Migration>}> </Route>
  </Routes>
  </Router>
  
   );
}

export default App;
