import { React, useState, useEffect } from "react";
import { config } from "../../config";
import { FaArrowAltCircleLeft, FaArrowAltCircleRight } from "react-icons/fa";
import { IconContext } from "react-icons";
import { data /*reemplazar data por bd*/, sliderSettings } from "../../data/CarouselData";
import { Row, Heading, Section, TextWrapper } from "../../globalStyles";
import {
    ButtonContainer,
    ReviewSlider,
    ImageWrapper,
    CarouselImage,
    CardButton,
} from "./CarouselStyles";
import {Routes, Route, useNavigate} from 'react-router-dom';

//agregar fetch to backend




const Carousel = () => {
    const [sliderRef, setSliderRef] = useState(null);
    const navigate = useNavigate();
    const [ data, setData ] = useState([]);

    useEffect( () => {
        getHotelData();
    }, [])

    const getHotelData = () => {
        fetch(config.HOST + ":" + config.PORT + "/hotels")
        .then(response => response.json())
        .then(data => {
            setData(data);
        });
    }


    const redirectLogic = (dir) => { //Redireccionar a pagina del hotel seleccionado
        navigate('/hotel/'+dir);
        console.log("Se redirigio a la página del hotel:" + dir)
    }

    return(
        <Section margin="0em auto 2em"  maxWidth="1280px" padding="150px, 70px, " inverse>
            <Row justify="cent" margin="1rem" wrap="wrap" paddingTop="2rem">
                <Heading inverse>
                    Paquetes y Promociones
                    {console.log(data)}
                </Heading>
                <ButtonContainer>
                    <IconContext.Provider value={{size:"4rem", color:"#48719b", }}>
                        <FaArrowAltCircleLeft onClick={sliderRef?.slickPrev}/>
                        <FaArrowAltCircleRight onClick={sliderRef?.slickNext}/>
                    </IconContext.Provider>
                </ButtonContainer>
            </Row>
            {data && data.length > 0 && (
  <ReviewSlider {...sliderSettings} ref={setSliderRef}>
    {data.map((data, index) => (
      <ImageWrapper key={index} onClick={() => redirectLogic(data.name)}>
        <CarouselImage src={data.image} />
        <TextWrapper size="1.1rem" margin="0.75rem 0 0">
          <strong style={{ display: "block", textAlign: "center" }}>{data.name}</strong>
          <br />
          {data.description}
        </TextWrapper>
        <CardButton> Mas Información </CardButton>
      </ImageWrapper>
    ))}
  </ReviewSlider>
)}
        </Section>
    );
}

export default Carousel;