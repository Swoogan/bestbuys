$(document).ready(function() {
  var focusoutEnabled = true;

  /* done */
  var focusout = function(name, command) {
    $("."+name).focusout(function(event) {      
      if (focusoutEnabled) { 
	var id = $(this).parents(".gameInfo").attr('data-game');
	var value = parseCurrency($(this).text());
	var data = '{"'+name+'": '+value+', "game": "'+id+'"}';
	post(command, data);
	blur(this);
      }
    });
  }

  /* done */
  var blur = function(element) {
    $(element).toggleClass("shadow");
    var newVal = currencyFormat(parseCurrency($(element).html()));
    $(element).html(newVal);
  }


  var reload = function() {
    var message = $('#message');
    message.removeClass('error');
    message.addClass('success');
    message.fadeIn('slow');
    message.html("Successfully saved changes.");

    var $tabs = $('#tabs').tabs();
    var sel = $tabs.tabs('option', 'selected')+1;
    load($('#tabs-'+sel));
  }

  var bindBehaviors = function() {
    /* done */
    $(".editable").focus(function() { 
      $(this).toggleClass("shadow");
    });

    $(".generate").click(function() { 
      var id = $(this).parents(".gameInfo").attr('data-game') ;
      var data = '{"game": "'+id+'"}';
      $.ajax({
              url: "/commands/", 
              type: "POST",
              data: '{"name":"generatePurchases", "data":'+data+'}',
              contentType: "application/json",
              dataType: "text",
              //success: reload
            }); 
    });

    focusout("income", "setIncome");
    focusout("upkeep", "setUpkeep");
    focusout("balance", "setBalance");
    focusout("wallet", "setWallet");
    focusout("landIncome", "setLandIncome");

    $(".structureCost").focusout(function(event) {
      if (focusoutEnabled) { 
        var id = $(this).parents(".gameInfo").attr('data-game') ;
        var value = parseCurrency($(this).text());
        var name = $(this).prev().text();
        var data = '{"structureCost": '+value+', "structureName": "'+name+'", "game": "'+id+'"}';
        post("setStructureCost", data);
        blur(this);
      }
    });

    /* done */
    $(".editable").keyup(function(e){ 
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

  var post = function(command, data) {
    $.ajax({
              url: "/commands/", 
              type: "POST",
              data: '{ "name": "'+command+'", "data": '+data+'}',
              contentType: "application/json",
              dataType: "text",
              success: reload
            }); 
  }

  /* done */
  var load = function(panel) {
    var id = $(panel).attr('data-game');
    $.getJSON('/games/'+id, function(data) {
      $.views.helpers({
        format: currencyFormat
      });
  
      $('.financeInfo').html(
        $('#financeTemplate').render(data)
      );

      $('.landsInfo').html(
        $('#landsTemplate').render(data['lands'])
      );

      $('.structuresInfo').html(
        $('#structuresTemplate').render(data['structures'])
      );
  
      $('.purchasesInfo').html(
        $('#purchasesTemplate').render(data['purchases'])
      );
  
      bindBehaviors();
    });
  }

  // This is sucky. Don't actually need to load the data after the first time.
  // tabscreate cause same data for all three, hmmmm...
  // commenting this out and the data still loads, looks like the page
  // initialization is the real problem.
  $("#tabs").on("tabsactivate", function(event, ui) {    
    load(ui.newPanel);
  });

  var url = '/games/?selector=%7B%22name%22%3A%201%7D'
  $.getJSON(url, function(data) {
    $('#tabs').html(
      '<ul>' + $('#tabItemTemplate').render(data) + '</ul>' + 
      $('#tabBodyTemplate').render(data)
    );

    /* n/a */
    $('.header').click(function() {
      $(this).next().slideToggle('slow');
      return false;
    }).next().hide();

    var $tabs = $('#tabs').tabs({
      show: function(e, ui) {
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

  /* done */
  $('#message').click(function(event) {
    $(this).fadeToggle(400);
  });

  /* done */
  $('#message').ajaxError(function(e, xhr, settings, exception) {
    $(this).removeClass('success');
    $(this).addClass('error');
    $(this).fadeToggle('slow');
    $('#error').text("Error in: '" + settings.url + "'");
    $('#exception').text('Message: ' + exception);
  });
});

