import 'dart:convert';

import 'package:holovn_mobile/models/schedule.dart';
import 'package:holovn_mobile/utils/config.dart';
import 'package:http/http.dart' as http;

Future<List<Schedule>> fetchSchedule() async {
  try {
    final scheduleResponse = await http.get(Uri.parse(readEnvironment("HOLOVN_SERVER") + "/schedules/current"));
    List scheduleList = json.decode(scheduleResponse.body);
    List<Schedule> values = List<Schedule>.from(scheduleList.map((schedule) => Schedule.fromJson(schedule)));
    values.sort((a,b) => a.startScheduled.compareTo(b.startScheduled));
    return values;
  } catch(err){
    print(err);
    return [];
  }
}