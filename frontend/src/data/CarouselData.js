export const data = [
    {
        title: "Hotel Ejemplo 1",
        image: "https://img.freepik.com/premium-vector/flat-hotel-building-background_23-2148146118.jpg",
        description: "¡Disfruta de unas vacaciones increíbles en este hotel!",
        rooms: "30",
    },
    {
        title: "Hotel Aleatorio",
        image: "https://img.freepik.com/premium-vector/flat-hotel-building-background_23-2148146118.jpg",
        description: "Descubre la comodidad y elegancia en este fantástico hotel.",
        rooms: "50",
    },
    {
        title: "Hotel de Ensueño",
        image: "https://img.freepik.com/premium-vector/flat-hotel-building-background_23-2148146118.jpg",
        description: "Experimenta un ambiente de lujo y tranquilidad en este hotel exclusivo.",
        rooms: "80",
    },
    {
        title: "Hotel Playa Paradise",
        image: "https://img.freepik.com/premium-vector/flat-hotel-building-background_23-2148146118.jpg",
        description: "Relájate en las hermosas playas cercanas a este encantador hotel.",
        rooms: "120",
    },
    {
        title: "Hotel Montaña Mágica",
        image: "https://img.freepik.com/premium-vector/flat-hotel-building-background_23-2148146118.jpg",
        description: "Admira las impresionantes vistas panorámicas de las montañas desde este hotel.",
        rooms: "40",
    },
    {
        title: "Hotel Del Mar",
        image: "https://img.freepik.com/premium-vector/flat-hotel-building-background_23-2148146118.jpg",
        description: "Sumérgete en un oasis de tranquilidad y confort en este hotel junto al mar.",
        rooms: "60",
    },
    {
        title: "Hotel Histórico",
        image: "https://img.freepik.com/premium-vector/flat-hotel-building-background_23-2148146118.jpg",
        description: "Explora la rica historia de la región mientras te hospedas en este hotel histórico.",
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
				slidesToShow: 2, // Cambiar a 2 máximo, POR DEFECTO 2
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