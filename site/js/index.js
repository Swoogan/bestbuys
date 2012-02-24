$(document).ready(function() {
  var focusoutEnabled = true;

  var blur = function(element) {
    $(element).toggleClass("shadow");
    var newVal = currencyFormat(parseCurrency($(element).html()));
    $(element).html(newVal);
  }

  var focusout = function(name, command) {
    $("."+name).focusout(function(event) {
      if (focusoutEnabled) { 
        var id = $(this).parents(".gameInfo").attr('data-game') 
        post(command, name, $(this).html(), id);
        blur(this);
      }
    });
  }

  var bindBehaviors = function() {
    $(".data").focus(function() { 
      $(this).toggleClass("shadow");
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
    var message = $('#message');
    message.removeClass('error');
    message.addClass('success');
    message.fadeToggle('slow');
    message.html("Successfully saved changes.");

    var $tabs = $('#tabs').tabs();
    var sel = $tabs.tabs('option', 'selected')+1;
    load($('#tabs-'+sel));
  }

  var post = function(command, name, value, game) {
    $.ajax({
              url: "/commands/", 
              type: "POST",
              data: '{ "name": "'+command+'", "data": {"'+name+'": '+parseCurrency(value)+', "game": "'+game+'"} }',
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
  
      $(panel).children(".financeInfo").html(
        $("#financeTemplate").render(data)
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

    $('.header').click(function() {
      $(this).next().slideToggle('slow');
      return false;
    }).next().hide();

    var $tabs = $('#tabs').tabs({
      show: function(e,ui){
        var $c = $(ui.panel), h = parseInt($(ui.tab).parent().outerHeight());
        while ( $c.length > 0 && $c.hasClass("gameInfo") ){
            h += ( $c.outerHeight(true) - $c.innerHeight() );
            $c = $c.parent();
        }
        h = parseInt($c.innerHeight() - h );
        $(ui.panel).height(h);
        return true;
      },
      add: function(e,ui){
          $(this).tabs("select",ui.index);
      }
    });
    var sel = $tabs.tabs('option', 'selected')+1;
    load($('#tabs-'+sel));
  });

  bindBehaviors();

  $('#message').click(function(event) {
    $(this).fadeToggle(400);
  });

  $('#message').ajaxError(function(e, xhr, settings, exception) {
    $(this).removeClass('success');
    $(this).addClass('error');
    $(this).fadeToggle('slow');
    $('#error').text("Error in: '" + settings.url + "'");
    $('#exception').text('Message: ' + exception);
  });
});

