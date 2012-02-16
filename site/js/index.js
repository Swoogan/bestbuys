$(document).ready(function() {
  var focusoutEnabled = true;

  var blur = function(element) {
    if (element.className.indexOf("shadow") != -1) {
      element.className = element.className.replace(" shadow", "");
      var newVal = currencyFormat(parseCurrency($(element).html()));
      $(element).html(newVal);
    }
  }

  var focusout = function(name, task) {
    $("."+name).focusout(function(event) {
      if (focusoutEnabled) { 
        post(task, name, $(this).html(), $(this).parents(".gameInfo").attr('data-game'));
        blur(this);
      }
    });
  }

  var bindBehaviors = function() {
    $(".data").focus(function() { 
      if (this.className.indexOf("shadow") == -1) {
        this.className += " shadow";
      }
    });

    focusout("income", "setIncome");
    focusout("upkeep", "setUpkeep");
    focusout("balance", "setBalance");
    focusout("wallet", "setWallet");
    focusout("landIncome", "setLandIncome");

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

  var post = function(task, name, value, game) {
    $.post( "/tasks/", 
            '{ "name": "'+task+'", "data": {"'+name+'": '+parseCurrency(value)+', "game": "'+game+'"} }', 
            reload);
  }

  var reload = function() {
     $.getJSON('/games/', function(data) {
       var listItems = "<ul>" + $("#listTemplate").render(data) + "</ul>";
       var games = $("#gameTemplate").render(data);

       $("#tabs").tabs("destroy");
       $("#tabs").html(listItems + games);
       $("#tabs").tabs();

       bindBehaviors();
    });
  }

  bindBehaviors();

  $.getJSON('/games/', function(data) {
    $.views.registerTags({
      format: currencyFormat
    });

    var listItems = "<ul>" + $("#listTemplate").render(data) + "</ul>";
    var games = $("#gameTemplate").render(data);

    $("#tabs").html(listItems + games);
    $("#tabs").tabs();

    bindBehaviors();
  });


  $("div.log").ajaxError(function(e, xhr, settings, exception) {
    $(this).text('error in: ' + settings.url + ' ' + exception);
  });
});

