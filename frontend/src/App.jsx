import React from 'react';
import GlobalStyle from './globalStyles';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';


// Components
import Navbar from './components/Navbar/Navbar'

// Pages
import Home from './pages/Home';
/* 
import Login from './pages/Login';
import Register from './pages/Register';
import Administrador from './pages/Administrador';
import Hotel from './pages/Hotel'; 
*/

function App() {
  return (
    <Router>
      <GlobalStyle/>
      <Navbar/>
      <Routes>
        <Route exact path='/' element={<Home/>} />
        {/* <Route path='/login' element={<Login/>} /> */}
        {/* <Route path='/register' element={<Register/>} /> */}
        {/* <Route path='/administrador' element={<Administrator/>} /> */}
        {/* <Route path='/hotel/:id' element={<Hotel/>} /> */}        
      </Routes>
    </Router>
  );
}

export default App;
