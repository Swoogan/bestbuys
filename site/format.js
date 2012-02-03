function CurrencyFormatted(amount)
{
  var f = parseFloat(amount);
  if(isNaN(f)) { return '$0'; }
  var s = new String(f);
  var temps = "";
  for (var i = s.length; i > 0; i -= 3)
  {
     temps = s.slice(Math.max(i-3, 0), i) + ',' + temps
  }

  return '$' + temps.slice(0, -1);
}
