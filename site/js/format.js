function currencyFormat(amount)
{
  var f = parseFloat(amount);
  if(isNaN(f)) { return "$0"; }

  var s = new String(f);
  var temps = "";
  for (var i = s.length; i > 0; i -= 3)
  {
     temps = s.slice(Math.max(i-3, 0), i) + "," + temps
  }

  return "$" + temps.slice(0, -1);
}

function parseCurrency(amount)
{
  if (amount.length === "") { return 0; }
  if (amount.charAt(0) == "$") { amount = amount.slice(1); }  
  amount = amount.replace(",", "");
  var i = parseInt(amount);
  if(isNaN(i)) { return "0"; }
  return i;
}
