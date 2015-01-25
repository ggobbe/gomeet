

$(function() {
  function initialize() {

    var mapOptions = {
      center: user.location,
      zoom: 10
    };

    var map = new google.maps.Map(document.getElementById('map-canvas'), mapOptions);
    var marker = new google.maps.Marker({
      position: user.location,
      map: map,
      title:user.name,
      icon: 'images/markers/darkGreen_MarkerA.png'
    });

    var removeAnimation = $.noop
    $.each(users, function(index, user){
      var marker = new google.maps.Marker({
        position: user.location,
        map: map,
        title: user.name,
        animation: google.maps.Animation.DROP
      });

      $("#usr-zoom-" + user.name).click(function(e){
        removeAnimation()
        $("h2")[0].scrollIntoView();
        map.setCenter(user.location);
        marker.setAnimation(google.maps.Animation.BOUNCE);
        removeAnimation = function(){ marker.setAnimation(null); }
        e.preventDefault();
      })
    });
  };

  initialize();
})
