$(document).ready(function() {
  var bindBehaviors = function() {
    $(".data").click(function() { this.contentEditable='true'; });

    $(".income").focusout(function() {
      $.post( "/tasks/", 
              '{ "name": "setIncome", "data": {"income": "' + $(this).html() + '"} }',
              {contentType: 'application/json'});
    });
    $(".upkeep").focusout(function() {
      $.post( "/tasks/", 
              '{ "name": "setUpkeep", "data": {"upkeep": "' + $(this).html() + '"} }',
              {contentType: 'application/json'});
    });
    $(".balance").focusout(function() {
      $.post( "/tasks/", 
              '{ "name": "setBalance", "data": {"balance": "' + $(this).html() + '"} }',
              {contentType: 'application/json'});
    });
    $(".wallet").focusout(function() {
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

