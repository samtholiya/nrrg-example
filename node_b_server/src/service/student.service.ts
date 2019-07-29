import { IServicePanel } from "../namespace/initializer";
import { Connection, Repository } from "typeorm";
import { Student } from "../entity/Student";
import { DatabaseService } from "./database.service";
import { MessagingService } from "./messaging.service";
import { RedisService } from "./redis.service";

@IServicePanel.register
export class StudentService {
    private static connection: Connection;
    private static repository: Repository<Student>;

    public init(): void {
        StudentService.connection = DatabaseService.getDatabaseConnection();
        StudentService.repository = StudentService.connection.getRepository(Student);
        MessagingService.listenToQueue("db.student", (message) => {
            console.log("Message received: " + message.getContent());
            RedisService.GetValueAndDelete(message.getContent(), (err, reply) => {
                let student: Student = JSON.parse(reply);
                StudentService.saveStudent(student).then((value)=>{
                    message.ack();
                });
            })
        })
    }

    public static saveStudent(role: Student): Promise<Student> {
        return this.repository.save(role);
    }

    public static getStudents(): Promise<Student[]> {
        return this.repository.find();
    }

}