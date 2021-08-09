import 'dart:convert';

import 'package:holovn_mobile/models/schedule_list.dart';
import 'package:http/http.dart' as http;

Future<ScheduleList> fetchSchedule() async {
  final scheduleResponse = await http.get(Uri.parse("https://localhost:8080/"));
  if (scheduleResponse.statusCode == 200){
    return ScheduleList.fromJson(json.decode(scheduleResponse.body));
  }
  throw Exception("Unable to fetch schedule");
}