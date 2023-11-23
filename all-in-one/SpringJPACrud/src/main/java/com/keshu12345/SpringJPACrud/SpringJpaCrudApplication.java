package com.keshu12345.SpringJPACrud;

import com.keshu12345.SpringJPACrud.DAO.StudentDAO;
import com.keshu12345.SpringJPACrud.entity.Student;
import org.springframework.boot.CommandLineRunner;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.context.annotation.Bean;

import java.util.List;

@SpringBootApplication
public class SpringJpaCrudApplication {

    public static void main(String[] args) {
        SpringApplication.run(SpringJpaCrudApplication.class, args);
    }
   @Bean
    public CommandLineRunner commandLineRunner(StudentDAO studentDAO) {
        return runner -> {
//            createStudent(studentDAO);
            createMultipleStudent(studentDAO);
//            readStudent(studentDAO);
//            queryForStudents(studentDAO);
//            queryForStudentByLastName(studentDAO);
//            updateStudent(studentDAO);
//            deleteStudent(studentDAO);
//            deleteAllStudent(studentDAO);

        };
    }

    private void deleteAllStudent(StudentDAO studentDAO) {
        System.out.println("Deleting all students");
        int numRowsDelete=studentDAO.deleteAll();
        System.out.println("delete row count: "+numRowsDelete);
    }

    private void deleteStudent(StudentDAO studentDAO) {
        int studentId=3002;
        studentDAO.delete(studentId);
        System.out.println("Deleted studentId : "+studentId);
    }

    private void updateStudent(StudentDAO studentDAO) {

        // retrieve student based on the id : primary key
        int theStudentId=3007;
        System.out.println("Getting student with studentId: "+theStudentId);
        Student myStudent=studentDAO.findById(theStudentId);
        // change  first name to "VivekanadRao"
        myStudent.setFirstName("VivekanadRao");
        myStudent.setLastName("Kankanti");

        // update the student
        studentDAO.update(myStudent);
        // display the updated student
        System.out.println("updated student:: "+myStudent);

    }

    private void queryForStudentByLastName(StudentDAO studentDAO) {
        // get list of student
        List<Student>theStudents=studentDAO.findByLastName("Rao");

        // display list of students

        for(Student theStudent:theStudents){
            System.out.println(theStudent);
        }
    }

    private void queryForStudents(StudentDAO studentDAO) {

        // get a list of students
        List<Student> theStudents=studentDAO.findAll();

        // display the list of student
        for(Student theStudent:theStudents){
            System.out.println(theStudent);
        }
    }

    private void readStudent(StudentDAO studentDAO) {
        // create a student object

        Student student=new Student("Keshav Spring","Rajput","spring@gmail.com");

        //save the object
        System.out.println("saved the student ...");
        studentDAO.save(student);

        // display id of the save student
         int theId=student.getId();
        System.out.println("saved student Generated id:: "+theId);

        // retrieve student based on the display primary key
        Student myStudent=studentDAO.findById(theId);

        // Display retrieve student
        System.out.println("found the student :: "+myStudent.getId()
                + " "+myStudent.getFirstName()
                +" "+myStudent.getLastName()
                +" "+myStudent.getEmail());
    }

    private void createMultipleStudent(StudentDAO studentDAO) {
        System.out.println("creating new student object... ");

        Student student1=new Student("Devendra","Kumar","Devendra@gmail.com");
        Student student2=new Student("Aman","Seth","Aman@gmail.com");
        Student student3=new Student("Vivekanad","Rao","vivek@gmail.com");


        // save the student object
        studentDAO.save(student1);
        studentDAO.save(student2);
        studentDAO.save(student3);
        System.out.println("saved student Generated id:: "+student1.getId());
        System.out.println("saved student Generated id:: "+student2.getId());
    }

    private void createStudent(StudentDAO studentDAO) {

        // create the student object
        System.out.println("creating new student object... ");

        Student student=new Student("Keshav","Rajput","onlyshafcs@gmail.com");

        // save the student object
        studentDAO.save(student);
        System.out.println("saved student Generated id:: "+student.getId());
    }
}

