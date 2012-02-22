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

  var reload = function() {
    var $tabs = $('#tabs').tabs();
    var sel = $tabs.tabs('option', 'selected')+1;
    load($('#tabs-'+sel));
  }

  var post = function(task, name, value, game) {
    $.ajax({
              url: "/tasks/", 
              type: "POST",
              data: '{ "name": "'+task+'", "data": {"'+name+'": '+parseCurrency(value)+', "game": "'+game+'"} }',
              contentType: "application/json",
              dataType: "text",
              success: reload
            }); 
  }

  var load = function(panel) {
    var id = $(panel).attr('data-game');
    $.getJSON('/games/'+id, function(data) {
      $.views.registerTags({
        format: currencyFormat
      });
  
      $(panel).html(
        $("#gameTemplate").render(data)
      );
  
      bindBehaviors();
    });
  }

  $('#tabs').bind('tabsselect', function(event, ui) {
    load(ui.panel);
  });

  var url = '/games/?selector=%7B%22name%22%3A%201%7D'
  $.getJSON(url, function(data) {
    $('#tabs').html(
      '<ul>' + $('#tabItemTemplate').render(data) + '</ul>' + 
      $('#tabBodyTemplate').render(data)
    );
    var $tabs = $('#tabs').tabs();
    var sel = $tabs.tabs('option', 'selected')+1;
    load($('#tabs-'+sel));
  });

  bindBehaviors();

  $("div.log").ajaxError(function(e, xhr, settings, exception) {
    $(this).text('error in: ' + settings.url + ' ' + exception);
  });
});

