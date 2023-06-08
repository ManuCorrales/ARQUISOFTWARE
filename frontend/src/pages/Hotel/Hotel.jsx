import { useEffect, useState } from 'react'; 
import {data} from '../../data/CarouselData'


function Hotel() {
    const [lastDirectory, setLastDirectory] = useState('');
    const [hotel, setHotel] = useState({});

    useEffect(() => {
        const pathArray = window.location.pathname.split('/');
        const lastPath = pathArray[pathArray.length - 1];
        setLastDirectory(lastPath);
        data.forEach(element => {

            if(decodeURIComponent(lastDirectory.replace(/\+/g, ' '))==element.title){setHotel(element)
 
            }})
      });
      console.log(hotel)



    return (
        <div style={{marginTop:"30px",width:"100vw",height:"100%",display:"flex",alignContent:"center",justifyContent:"center"}}>
            <div style={{marginTop:"30px",width:"900px",height:"1000px", padding:"40px", borderRadius:"10px", backgroundColor:"white"}}>
                <div style={{width:"100%",fontSize:"40px"}}>{hotel.title}</div>
                <div style={{padding:"10px"}}>{hotel.description}</div>
                <img src={hotel.image} style={{width:"100%"}}/>
            </div>
        </div>
    );
  }
  export default Hotel
