
export const data = [
    {
        name: "Hotel Transylvania",
        image: "https://media.lacapital.com.ar/p/0170d1bbb6dc9efd3603950594e87d15/adjuntos/203/imagenes/007/401/0007401921/642x0/smart/el-hotel-transylvania-donde-los-monstruos-estan-como-su-casa-y-lejos-sus-horrorosas-pesadilla-los-humanos.jpg",
        description: "¡Disfruta de unas vacaciones asombrosas mientras tus monstruos favoritos te cuentan sus mejores historias!",
        rooms: 30,
        email: "hotel@transylvania.spook",
        telephone: "+66 666-666",
    },
    {
        name: "Hotel Continental",
        image: "https://live.staticflickr.com/65535/52129719199_c868c4951d_o.jpg",
        description: "El legendario hotel para los asesinos más picantes del mundo, la discresion y los codigos son sagrados, no querras encontrarte a John Wick.",
        rooms: "50",
    },
    {
        name: "Caesars Palace",
        image: "https://i.ytimg.com/vi/PXwcp2LymEc/maxresdefault.jpg",
        description: "No olvidaras estas experiencias de desenfreno, fiesta y locura. Cuando te vayas diras: ¿Qué Pasó Ayer?.",
        rooms: "80",
    },
    {
        name: "THE GOAT",
        image: "https://s3.amazonaws.com/arc-wordpress-client-uploads/infobae-wp/wp-content/uploads/2018/03/29155730/Hotel-Messi-Ibiza-Es-Vive-10-1920.jpg",
        description: "Disfruta una estadía inolvidable en el hotel del jugador mas famoso del mundo, eso si, no olvides tu billetera...",
        rooms: "120",
    },
    {
        name: "Mandalay Bay",
        image: "https://mandalaybay.mgmresorts.com/content/dam/MGM/mandalay-bay/hero-shots/mandalay-bay-architecture-hero-shot.tif",
        description: "Lugar donde los sueños se convierten en realidad y cada día es una oportunidad para superar tus límites y alcanzar la victoria, como Rocky lo hizo.",
        rooms: "40",
    },
    {
        name: "Mansion Foster",
        image: "https://uvn-brightspot.s3.amazonaws.com/assets/vixes/f/foster_home.jpg",
        description: "Descubre habitaciones donde tus amigos imaginarios cobraran vida y la locura es bienvenida.",
        rooms: "60",
    },
    {
        name: "Mi Pobre Angelito",
        image: "https://cloudfront-us-east-1.images.arcpublishing.com/coxohio/6534CBMDHVGUBV5OONX6XHDEEU.jpg",
        description: "La adrenalina y la accion no tienen limites, fijate bien donde pisas, que tocas y que respiras, hay trampas por doquier, necesitaras tus 5 sentidos muy despiertos.",
        rooms: "20",
    },
    {
        name: "El Gran Hotel Budapest",
        image: "https://images.adsttc.com/media/images/5318/9a36/c07a/806c/d900/0134/large_jpg/20140302-GRAND-slide-SOK2-superJumbo.jpg?1394121259",
        description: "Quieres vivir una historia donde la gravedad es desafiada, y no tienes miedo a perderte? ¡Prepárate para un viaje cómico que nunca olvidarás!.",
        rooms: "20",
    },
];

export const sliderSettings = {
	arrows: false,
	slidesToShow: 3, //Cambiar a 3 maximo, si hay menos novedades poner ese numero.
	focusOnselect: false,
	accessability: false,
	responsive: [
		{
			breakpoint: 1280,
			settings: {
				slidesToShow: 1, // Cambiar a 2 máximo, POR DEFECTO 2
			},
		},
		{
			breakpoint: 900,
			settings: {
				slidesToShow: 1, // NO EDITAR, POR DEFECTO 1
			},
		},
	],
};