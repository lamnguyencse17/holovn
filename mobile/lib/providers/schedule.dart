import 'dart:convert';

import 'package:holovn_mobile/models/schedule.dart';
import 'package:http/http.dart' as http;

Future<List<Schedule>> fetchSchedule() async {
  try {
    final scheduleResponse = await http.get(Uri.parse("http://localhost:8080/schedules/current"));
    List scheduleList = json.decode(scheduleResponse.body);
    List<Schedule> values = List<Schedule>.from(scheduleList.map((schedule) => Schedule.fromJson(schedule)));
    return values;
  } catch(err){
    print(err);
    throw Exception("Unable to fetch schedule");
  }

}