'use strict'


let today = new Date();
let formatDate = today.toDateString();
let selectElement = document.getElementById('date');
selectElement.innerHTML = formatDate;
//console.log('Here\'s a hidden message');
var map;
 function risk(number){
      if (number < 5) {
        return "низький";
      } else if (number > 4 && number < 13) {
        return "середній";
      } else if (number > 12) {
        return "високий";
      }
 }
function initMap() {
    var kyiv = {lat: 50.434341, lng: 30.527756};
    map = new google.maps.Map(document.getElementById('map'), {
    center: kyiv,
    zoom: 4
  });
  
  var infowindow = new google.maps.InfoWindow();
  var marker;
  var markers = [];
  var bounds = new google.maps.LatLngBounds();
    $.getJSON("test.json", function(data) {
      $.each(data.Places, function(key, place) {
        var latLng = new google.maps.LatLng(place[0].Lat, place[0].Lng); 
        alert(latlng);
        // Creating a marker and putting it on the map
        var marker = new google.maps.Marker({
            position: latLng,
            map: map,
            title: place[0].Name
        });
      markers.push(marker);
      bounds.extend(marker.position);
      google.maps.event.addListener(marker, "click", (function (mm, tt) { //can be changed with 'click' event
            return function () {                
                var infoContent = '<div class="infowindow">';
                infoContent += '<div class="point-name">'+ place[0].Name+'</div>';
                infoContent += '<div class="point-address">' + risk(place['Intensity']) + '</div>';
                infoContent += "</div>";

                infowindow.setOptions({
                    content: infoContent
                });

                infowindow.open(map, mm);
            };
        })(marker, place[0].Name));
    });
    map.fitBounds(bounds);
  // Try HTML5 geolocation.
  if (navigator.geolocation) {
    navigator.geolocation.getCurrentPosition(function(position) {
      var pos = {
        lat: position.coords.latitude,
        lng: position.coords.longitude
      };

      infoWindow.setPosition(pos);
      //infoWindow.setContent('Location found.');
      //infoWindow.open(map);
      map.setCenter(pos);
    }, function() {
      handleLocationError(true, infoWindow, map.getCenter());
    });
  } else {
    // Browser doesn't support Geolocation
    handleLocationError(false, infoWindow, map.getCenter());
  }
});

function handleLocationError(browserHasGeolocation, infoWindow, pos) {
  infoWindow.setPosition(pos);
  infoWindow.setContent(browserHasGeolocation ?
                        'Error: The Geolocation service failed.' :
                        'Error: Your browser doesn\'t support geolocation.');
  infoWindow.open(map);
}
// Check for the various File API support.
if (window.File && window.FileReader && window.FileList && window.Blob) {
  // Great success! All the File APIs are supported.
} else {
  alert('The File APIs are not fully supported in this browser.');
}

function handleFileSelect(evt) {
  var files = evt.target.files; // FileList object

  // files is a FileList of File objects. List some properties.
  var output = [];
  for (var i = 0, f; f = files[i]; i++) {
    output.push('<li><strong>', escape(f.name), '</strong> (', f.type || 'n/a', ') - ',
                f.size, ' bytes, last modified: ',
                f.lastModifiedDate.toLocaleDateString(), '</li>');
  }
  document.getElementById('list').innerHTML = '<ul>' + output.join('') + '</ul>';
}

document.getElementById('files').addEventListener('change', handleFileSelect, false);