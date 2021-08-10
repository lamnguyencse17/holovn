import 'dart:convert';
import 'package:flutter_dotenv/flutter_dotenv.dart';
import 'package:holovn_mobile/models/schedule.dart';
import 'package:http/http.dart' as http;

Future<List<Schedule>> fetchSchedule() async {
  try {
    print(dotenv.env["HOLOVN_SERVER"]);
    final scheduleResponse = await http.get(Uri.parse(dotenv.env["HOLOVN_SERVER"]! + "/schedules/current"));
    List scheduleList = json.decode(scheduleResponse.body);
    List<Schedule> values = List<Schedule>.from(scheduleList.map((schedule) => Schedule.fromJson(schedule)));
    return values;
  } catch(err){
    print(err);
    return [];
  }

}