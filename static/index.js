$(function() {
  var uploadButton = $('#uploadButton');
  uploadButton.on('click', function() {
    $.get('/api/setImageReady')
    .done(function(data) {
      $('#successModal .modal-body').text(data.rowsUpdated + " new images set ready to upload.", null, 2);  
      var modal = $('#successModal').modal('show');
      setTimeout(function() {
        modal.modal('hide'); 
      }, 3000);
    })
    .fail(function(jqXHR, textStatus, errorThrown) {
      $('#failureModal .modal-body').text(JSON.stringify(jqXHR, null, 2));
      var modal = $('#failureModal').modal('show');
      setTimeout(function() {
        modal.modal('hide'); 
      }, 3000);
    });
    
  }); 
});

// Click handler for rows
$('.tr-image-rows').click(function() {

    // Get ID value
    var id = $(this).attr('id');
  
    // Split to get image filename 
    var imageFname = id.split('-')[1];
  
    // Redirect 
    window.location.href = '/imageDetail?imageFname=' + imageFname;

  });