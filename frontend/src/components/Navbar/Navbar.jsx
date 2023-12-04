import { data } from '../../data/NavbarData.jsx'
import { React, useState, useEffect } from 'react';
import { FaTimes } from 'react-icons/fa';
import { CgMenuRight } from 'react-icons/cg';
import { IconContext } from 'react-icons';
import "./Navbar.css";
import {
    Nav, 
    NavbarContainer, 
    NavLogo, 
    NavIcon, 
    MobileIcon,
    NavMenu,
    NavItem,
    NavLinks
} from './NavbarStyles.jsx';
import { Routes, Route, useLocation, useNavigate } from 'react-router-dom';



const Navbar = () => {
    let config = {
        HOST: "http://localhost",
        PORT: 8090
    }
    const [show, setShow] = useState(false);

    let navigate = useNavigate();
    let location = useLocation();



    const handleClick = () => {
        if (show === true){
            document.querySelector("#nav-container").style.position="relative";
            
        }
        else{
            document.querySelector("#nav-container").style.position="sticky";
            
        }

        setShow(!show);
    }

    const closeMobileMenu = (to, text) => {

        // No funciona
        /*if (text && location.pathname==='/') {
            scrollTo(text);
          }*/

        document.querySelector("#nav-container").style.position="relative";
        
        navigate(to);
        
        setShow(false);
    }


    const scrollTo = (id) => {
        const element = document.getElementById(id)

        element.scrollIntoView({
            behavior: 'smooth',
        })
    }

    
    var Administrador = false

    return(
        <IconContext.Provider value={{color:'black'}}>
            <Nav id="nav-container">
                <NavbarContainer>
                    <NavLogo to="/">  {/*me manda al home*/}
                        <NavIcon src="../assets/logo.png" alt="logo" style={{ width: '111px', height: '111px'}}/>
                    </NavLogo>
                    <MobileIcon onClick={handleClick}>
                        {show ? <FaTimes/> : <CgMenuRight/>}
                    </MobileIcon>
                    <NavMenu show={show}>
                        {data.filter(item => item.text !== "Administrador" || Administrador) // sera true siermpre que no sea administrador ,
                             .map((item, index) => (
                            <NavItem key={index}>
                                <NavLinks onClick={() => closeMobileMenu(item.to, item.text)}> {item.text} </NavLinks>
                            </NavItem>
                        ))}
                        <NavItem key={5}>
                                <NavLinks onClick={() => closeMobileMenu("/listaReservas", "Reservas")}> Reservas </NavLinks>
                            </NavItem>
                    </NavMenu>
                </NavbarContainer>
            </Nav>
        </IconContext.Provider>
    );
}

export default Navbar;