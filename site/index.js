$(document).ready(function() {
  var bindBehaviors = function() {
    $(".data").click(function() { 
      if (this.className.indexOf("shadow") == -1) {
        this.className += " shadow";
      }
    });

    $(".income").focusout(function() {
      $.post( "/tasks/", 
              '{ "name": "setIncome", "data": {"income": "' + $(this).html() + '"} }',
              {contentType: 'application/json'});
      if (this.className.indexOf("shadow") != -1) {
        this.className = this.className.replace(" shadow", "");
      }
    });
    $(".upkeep").focusout(function() {
      $.post( "/tasks/", 
              '{ "name": "setUpkeep", "data": {"upkeep": "' + $(this).html() + '"} }',
              {contentType: 'application/json'});
      if (this.className.indexOf("shadow") != -1) {
        this.className = this.className.replace(" shadow", "");
      }
    });
    $(".balance").focusout(function() {
      $.post( "/tasks/", 
              '{ "name": "setBalance", "data": {"balance": "' + $(this).html() + '"} }',
              {contentType: 'application/json'});
      if (this.className.indexOf("shadow") != -1) {
        this.className = this.className.replace(" shadow", "");
      }
    });
    $(".wallet").focusout(function() {
      $.post( "/tasks/", 
              '{ "name": "setWallet", "data": {"wallet": "' + $(this).html() + '"} }',
              {contentType: 'application/json'});
      if (this.className.indexOf("shadow") != -1) {
        this.className = this.className.replace(" shadow", "");
      }
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

