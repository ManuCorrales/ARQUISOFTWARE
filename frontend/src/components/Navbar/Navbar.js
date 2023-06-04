import { React, useState } from 'react';
import { FaTimes } from 'react-icons/fa';
import { CgMenuRight } from 'react-icons/cg';
import { IconContext } from 'react-icons';
import {
    Nav, 
    NavbarContainer, 
    NavLogo, 
    NavIcon, 
    MobileIcon,
    NavMenu,
    NavItem,
    NavLinks
} from './NavbarStyles.js';
import { useLocation, useNavigate } from 'react-router-dom';
import { data } from '../../data/NavbarData';

const Navbar = () => {
    const [show, setShow] = useState(false);

    let navigate = useNavigate();
    let location = useLocation();

    const handleClick = () => {
        if (show === true){
            document.querySelector("#nav-container").style.position="absolute";
            document.querySelector("#nav-container").style.backgroundColor="limegreen";
        }
        else{
            document.querySelector("#nav-container").style.position="sticky";
            document.querySelector("#nav-container").style.backgroundColor="green";
        }

        setShow(!show);
    }

    const closeMobileMenu = (to, text) => {


        if (text && location.pathname==='/') {
            scrollTo(text);
        }

        document.querySelector("#nav-container").style.position="absolute";
        document.querySelector("#nav-container").style.backgroundColor="limegreen";
        navigate(to);
        
        setShow(false);
    }


    const scrollTo = (id) => {
        const element = document.getElementById(id)

        element.scrollIntoView({
            behavior: 'smooth',
        })
    }

    return(
        <IconContext.Provider value={{color:'black'}}>
            <Nav id="nav-container">
                <NavbarContainer>
                    <NavLogo to="/">
                        <NavIcon src="./assets/logo.png" alt="logo"/>
                    </NavLogo>
                    <MobileIcon onClick={handleClick}>
                        {show ? <FaTimes/> : <CgMenuRight/>}
                    </MobileIcon>
                    <NavMenu show={show}>
                        {data.map((item, index) => (
                            <NavItem key={index}>
                                <NavLinks onClick={() => closeMobileMenu(item.to, item.text)}> {item.text} </NavLinks>
                            </NavItem>
                        ))}
                    </NavMenu>
                </NavbarContainer>
            </Nav>
        </IconContext.Provider>
    );
}

export default Navbar;