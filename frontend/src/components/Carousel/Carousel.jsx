import { React, useState } from "react";
import { FaArrowAltCircleLeft, FaArrowAltCircleRight } from "react-icons/fa";
import { IconContext } from "react-icons";
import { data /*reemplazar data pro bd*/, sliderSettings } from "../../data/CarouselData";
import { Row, Heading, Section, TextWrapper } from "../../globalStyles";
import {
    ButtonContainer,
    ReviewSlider,
    ImageWrapper,
    CarouselImage,
    CardButton,
} from "./CarouselStyles";

//agregar fetch to backend

const Carousel = () => {
    const [sliderRef, setSliderRef] = useState(null);

    /*
    //agregar control de datos de la bd, (crear array de hoteles desde la bd)
    const [ data, setData ] = useState([]);
    const [ needData, setNeedData ] = useState(true);

    if (needData){
        //fetch db fill data array with setData
    }*/

    


    const redirectLogic = (dir) => { //Redireccionar a pagina del hotel seleccionado
        console.log("Se redirigio a la página del hotel:" + dir)
    }

    return(
        <Section margin="0em auto 2em"  maxWidth="1280px" padding="150px, 70px, " inverse>
            <Row justify="cent" margin="1rem" wrap="wrap" paddingTop="2rem">
                <Heading inverse>
                    Paquetes y Promociones
                </Heading>
                <ButtonContainer>
                    <IconContext.Provider value={{size:"4rem", color:"#48719b", }}>
                        <FaArrowAltCircleLeft onClick={sliderRef?.slickPrev}/>
                        <FaArrowAltCircleRight onClick={sliderRef?.slickNext}/>
                    </IconContext.Provider>
                </ButtonContainer>
            </Row>
            <ReviewSlider {...sliderSettings} ref={setSliderRef}>
                {data.map((data, index) => (
                    <ImageWrapper key={index} onClick={() => redirectLogic(data.title)}>
                        <CarouselImage src={data.image}/>
                        <TextWrapper size="1.1rem" margin="0.75rem 0 0" >
                            <strong style={{ display: "block", textAlign: "center" }}>{data.title}</strong>
                            <br/>
                            {data.description}
                        </TextWrapper>
                        <CardButton> Mas Información </CardButton>
                    </ImageWrapper>
                ))}
            </ReviewSlider>
        </Section>
    );
}

export default Carousel;