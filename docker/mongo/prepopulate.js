// Set data connection values
var db = db.getSiblingDB("alerts_db");
var alertsCollection = db.getCollection("alerts");

// Generate random data for alerts docs
function generateRandomAlert() {
  var names = ["Alert 1", "Alert 2", "Alert 3", "Alert 4", "Alert 5"];
  var types = ["Type A", "Type B", "Type C", "Type D", "Type E"];
  var messages = ["Message 1", "Message 2", "Message 3", "Message 4", "Message 5"];

  var randomIndex = Math.floor(Math.random() * names.length);

  return {
    name: names[randomIndex],
    type: types[randomIndex],
    message: messages[randomIndex],
  };
}

// Insert docs for alerts
for (var i = 0; i < 20; i++) {
  var alert = generateRandomAlert();
  alertsCollection.insert(alert);
}