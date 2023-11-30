import Register from "../../pages/register/Register";
import { useEffect, useState } from 'react'; 
import Login from "../../pages/login/login.jsx";

const Modal = ({funcToggle}) => {
    const [registrar, setRegistrar] = useState(false)
    const [loginToggle, setLogin] = useState(false)

    return(
        <div style={{left:"0", top:"0",zIndex:"49",width:"100%",height:"100%",position:"fixed",backgroundColor:"rgba(0,0,0,0.5)",alignContent:"center",justifyContent:"center",display:"flex"}}>{/*modal background*/}
            <div style={{
                backgroundColor:"white",
                borderRadius:"10px",
                width:"400px",
                top:"100px",
                alignItems:"stretch",
                marginTop:"100px",
                position:"absolute",
                alignContent:"center",
                justifyContent:"center",
                padding:"10px",
                display:"flex",
                flexDirection: "column"}}>{/*modal main*/}
                <div style={{fontSize:"25px",width:"100%",alignContent:"center",justifyContent:"center"}}>{/*modal titulo*/}Quien eres?</div>
                <div style={{fontSize:"15px",alignContent:"center", justifyContent:"center"}}>{/*modal subtitulo*/}Debes registrarte para hacer una reserva</div>
                <button onClick={() => {setRegistrar(!registrar);setLogin(false)}}>{/*modal boton registro*/}Registrarse</button>
                <button onClick={() => {setLogin(!loginToggle);setRegistrar(false)}}>{/*modal boton registro*/}Iniciar Sesion</button>
                {registrar&&<Register home={false}/>}
                {loginToggle&&<Login home={false}/>}
                <button style={{ position:"absolute",top:"0px", right:"0px"}} onClick={funcToggle}>{/*modal boton salir*/}X</button>
            </div>
        </div>


    );
}

export default Modal;