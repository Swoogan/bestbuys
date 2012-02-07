$(document).ready(function() {
  var bindBehaviors = function() {
    $(".data").click(function() { this.contentEditable='true'; });

    $(".data").focusout(function() {
      alert($(this).html());
      $.post( "/tasks/", 
              '{ "name": "setWallet", "data": {"wallet": "' + $(this).html() + '"} }',
              {contentType: 'application/json'});
    });

    $(".data").keypress(function(e){ return e.which != 13; });
  }

  bindBehaviors();

  $.getJSON('/games/', function(data) {
     $.views.registerTags({
       format: CurrencyFormatted
     });

     $("#gameList" ).html(
       $("#gameTemplate").render(data)
     );
     bindBehaviors();
  });

  $("div.log").ajaxError(function(e, xhr, settings, exception) {
    $(this).text('error in: ' + settings.url + '  error:' + xhr.responseText);
  });
});

