package com.keshu12345.SpringJPACrud.DAO;

import com.keshu12345.SpringJPACrud.entity.Student;

import java.util.List;

public interface StudentDAO {

    void save(Student student);
    Student findById(Integer id);
    List<Student> findAll();
    List<Student>findByLastName(String theLastName);
    void update(Student student);
    void delete(Integer studentId);
    int deleteAll();
}
