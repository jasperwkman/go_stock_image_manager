$(document).ready(function() {

    function flashInput(elementId, color) {
        const $input = $("#"+elementId);
        
        if (color === 'green') {
          $input.css('backgroundColor', '#d5f5e3');
          $input.css('background-color', '#d5f5e3');
        } else if (color === 'yellow') {  
            $input.css('backgroundColor', '#f9e79f');
            $input.css('background-color', '#f9e79f');
        } else if (color === 'red') {
          $input.css('backgroundColor', '#f5b7b1'); 
          $input.css('background-color', '#f5b7b1'); 
        }
      
        setTimeout(() => {
          $input.css('backgroundColor', 'white'); 
          $input.css('background-color', 'white'); 
        }, 3000);
      }

    // The DOM element you wish to replace with Tagify
    var input = document.querySelector('input[id=in_imageTag]');

    // initialize Tagify on the above input node reference
    var tagify = new Tagify(input, {
        originalInputValueFormat: valuesArr => valuesArr.map(item => item.value).join(',')
    });
    tagify.on('remove', function(e) {
        console.log(e.detail);
    });

    $('.legend-tags').on('click', function() {
        var id = $(this).attr('id');
        tagify.addTags('#TAGS-'+id.split('-')[1])
      });

    $('#addTagButton').on('click', function() {
        tagify.addTags(['abc']);
    });
    $('#btn-update-tags').on('click', function() {
            // Get the image filename and tags from the page
            var imageFname = $("#td_imageFname").text();
            var imageTags = tagify.value.map(item => item.value).join(',');

            // Create the data object to send to the server
            var data = {
                "imageFname": imageFname, 
                "imageTags": imageTags
            };
            
            console.log(data);
            // Send AJAX request to update tags
            $.ajax({
                url: "/api/setImageTags", 
                method: "POST",
                contentType: "application/json",
                data: JSON.stringify(data),
                success: function(response) {
                    console.log(response);
                    // This function will be called when the request is successful
                    if (response['rowsUpdated'] == 1) {
                        flashInput("td_imageTag","green")
                    }else{
                        flashInput("td_imageTag","yellow")
                    }
                    if (response['error'] != "no"){
                        flashInput("td_imageTag","red")
                    }
                },

                error: function(jqXHR, textStatus, errorThrown) {
                    // This function will be called when an error occurs during the request
                    console.error(textStatus);
                    console.error(errorThrown);
                    flashInput("td_imageTag","red");
                }

        });
    });

    // This function creates an on/off switch for a stock image provider
    function createOnOffSwitch(inputId, label, statusId) {
        // Create the on/off switch element
        var onOffSwitch = $(`<label class="switch"><input type="checkbox" id="${statusId}-switch"><span class="slider round"></span></label>`);
        // If the value of the input element is 1, set the switch to checked
        if (parseInt($(inputId).val()) === 1) {
            onOffSwitch.find('input').prop('checked', true);
        }
        // Listen for changes to the switch
        onOffSwitch.find('input').on('change', function() {
            console.log("statusId: " + statusId);
            // Get whether the switch is checked or not
            var checked = $(this).prop('checked');
            // Update the value of the input element
            $(inputId).val(checked ? 1 : 0);
            // Update the text next to the switch
            $(`#${statusId}-text`).text(checked ? 'upload on' : 'upload off');
            // Set the statusValue based on whether the switch is checked or not
            var statusValue = checked ? 1 : 0;
            // Create the data object to send to the server
            var data = {
                "imageFname": $("#td_imageFname").text(),
                "stock": statusId.split("-")[0],
                "statusValue": statusValue
            };
            console.log(data);
            // Send a POST request to the server with the data
            res = $.ajax({
                url: "/api/setImageStatus",
                method: "POST",
                contentType: "application/json",
                data: JSON.stringify(data),
                success: function(response) {
                    // This function will be called when the request is successful
                    if (response['rowsUpdated'] == 1) {
                        flashInput(statusId,"green")
                    }else{
                        flashInput(statusId,"yellow")
                    }

                    if (response['error'] != "no"){
                        flashInput(statusId,"red")
                    }
                },
                error: function(jqXHR, textStatus, errorThrown) {
                    // This function will be called when an error occurs during the request
                    flashInput(statusId,"red")
                }
            });
        });
        // Append the switch and text to the label element
        label.append(onOffSwitch);
        label.append(`<span class="on-off-text" id="${statusId}-text">${parseInt($(inputId).val()) === 1 ? 'upload on' : 'upload off'}</span>`);
    }

    // This function creates a tag label for a stock image provider
    function createTagLabel(inputId, label, statusId) {
        // Create the tag label element
        var tagLabel = $(`<span class="tag" id="${statusId}-tag"></span>`);
        // Get the value of the input element
        var value = parseInt($(inputId).val());
        // Set the text and class of the tag label based on the value
        if (value === 10) {
            tagLabel.addClass('processing');
            tagLabel.text('Processing');
        } else if (value === 11) {
            tagLabel.addClass('completed');
            tagLabel.text('Completed');
        } else {
            tagLabel.addClass('unknown');
            tagLabel.text('UNKNOWN');
        }
        // Append the tag label to the label element
        label.append(tagLabel);
    }

    // This function updates the status of a stock image provider
    function updateStatus(inputId, label, statusId) {
        // Get the current value of the input element
        var value = parseInt($(inputId).val());
        // If the value is 0 or 1, create an on/off switch
        if (value === 0 || value === 1) {
            createOnOffSwitch(inputId, label, statusId);
        } else {
            // If the value is not 0 or 1, create a tag label
            createTagLabel(inputId, label, statusId);
        }
    }

    // Listen for clicks on the "Update" button
    $('.update-description').on('click', function() {
        // Get the image filename and description from the page
        var imageFname = $("#td_imageFname").text();
        var imageDescription = $("#txt-imageDescription").val();

        // Create the data object to send to the server
        var data = {
            "imageFname": imageFname,
            "imageDescription": imageDescription
        };

        // Send a POST request to the server with the data
        $.ajax({
            url: "/api/setImageDescription",
            method: "POST",
            contentType: "application/json",
            data: JSON.stringify(data),
            success: function(response) {
                // This function will be called when the request is successful
                if (response['rowsUpdated'] == 1) {
                    flashInput("txt-imageDescription","green")
                }else{
                    flashInput("txt-imageDescription","yellow")
                }

                if (response['error'] != "no"){
                    flashInput("txt-imageDescription","red")
                }
            },
            error: function(jqXHR, textStatus, errorThrown) {
                // This function will be called when an error occurs during the request
                flashInput("txt-imageDescription","red")
            }
        });
    });

    // Update the status of each stock image provider
    updateStatus('#foap-status-input', $('#foap-status .status-label'), 'foap-status');
    updateStatus('#shutterstock-status-input', $('#shutterstock-status .status-label'), 'shutterstock-status');
    updateStatus('#alamy-status-input', $('#alamy-status .status-label'), 'alamy-status');

});
