import React from 'react';
import GlobalStyle from './globalStyles';

import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';


// Components
import Navbar from './components/Navbar/Navbar';

// Pages
import Home from './pages/home/Home';
import Login from './pages/login/login';

import Register from './pages/register/Register';
/* 
import Administrador from './pages/Administrador';*/
import Hotel from './pages/Hotel/Hotel'; 
import Admin from './pages/admin/Admin';
import ListaReservas from './pages/listaReservas/ListaReservas';

function App() {
  return (
    <Router>
      <GlobalStyle/>
      <Navbar/>
      <Routes>
        <Route exact path='/' element={<Home />} />
        <Route exact path='/admin' element={<Admin />} />
        <Route exact path='/listaReservas' element={<ListaReservas />} />
        <Route path='/login' element={<Login home={true}/>} /> 
        <Route path='/register' element={<Register home={true}/>} /> 
        {/* <Route path='/administrador' element={<Administrator/>} /> */}
        <Route path='/hotel/:id' element={<Hotel/>} />        
      </Routes>
    </Router>
  );
}

export default App;