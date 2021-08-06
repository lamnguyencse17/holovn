import 'dart:convert';

import 'package:holovn_mobile/models/schedule.dart';
import 'package:http/http.dart' as http;

Future<Schedule> fetchSchedule() async {
  final scheduleResponse = await http.get(Uri.parse("https://localhost:8080/"));
  if (scheduleResponse.statusCode == 200){
    return Schedule.fromJson(json.decode(scheduleResponse.body));
  }
  throw Exception("Unable to fetch schedule");
}