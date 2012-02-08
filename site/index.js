$(document).ready(function() {
  var focusoutEnabled = true;

  var blur = function(element) {
    if (element.className.indexOf("shadow") != -1) {
      element.className = element.className.replace(" shadow", "");
    }
  }

  var bindBehaviors = function() {
    $(".data").click(function() { 
      if (this.className.indexOf("shadow") == -1) {
        this.className += " shadow";
      }
    });

    $(".income").focusout(function() {
      if (!focusoutEnabled) { return; }
      post("setIncome", "income", $(this).html());
      blur(this);
    });
    $(".upkeep").focusout(function() {
      if (!focusoutEnabled) { return; }
      post("setUpkeep", "upkeep", $(this).html());
      blur(this);
    });
    $(".balance").focusout(function() {
      if (!focusoutEnabled) { return; }
      post("setBalance", "balance", $(this).html());
      blur(this);
    });
    $(".wallet").focusout(function() {
      if (!focusoutEnabled) { return; }
      post("setWallet", "wallet", $(this).html());
      blur(this);
    });

    $(".data").keyup(function(e){ 
      var esc = e.which == 27;
      var enter = e.which == 13;

      if (esc) {
        document.execCommand("undo", false, null);
        focusoutEnabled = false;
        e.target.blur();
        blur(e.target);
        focusoutEnabled = true;
      } else if (enter) {
        e.target.blur();
      }
    });
  }

  var post = function(task, name, value) {
    $.post( "/tasks/", 
            '{ "name": "'+task+'", "data": {"'+name+'": "'+ParseCurrency(value)+'"} }',
            {contentType: 'application/json'});
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

