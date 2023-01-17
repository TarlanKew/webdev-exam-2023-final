let map = L.map('map', {
    layers: MQ.hybridLayer(),
    center: [55.75394873322216, 37.62079491226199],
    zoom: 16
});

L.marker(
    [55.75394873322216, 37.62079491226199]
).addTo(map);

document.getElementsByClassName("leaflet-bottom leaflet-left")[0].remove();
document.getElementsByClassName("leaflet-bottom leaflet-right")[0].remove();