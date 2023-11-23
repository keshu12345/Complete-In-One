package com.keshu12345.SpringJPACrud.DAO;

import com.keshu12345.SpringJPACrud.entity.Student;
import jakarta.persistence.EntityManager;
import jakarta.persistence.TypedQuery;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Repository;
import org.springframework.transaction.annotation.Transactional;

import java.util.List;

@Repository
public class StudentDAOImpl implements StudentDAO{
    // Define field for entity manager

    private EntityManager entityManager;

    // Inject the dependency in constructor
    @Autowired  // this is optional if you have only one constructor
    public StudentDAOImpl(EntityManager entityManager) {
        this.entityManager = entityManager;
    }
    @Transactional
    @Override
    public void save(Student student) {
        entityManager.persist(student);

    }

    @Override
    public Student findById(Integer id) {
        return entityManager.find(Student.class,id);
    }

    @Override
    public List<Student> findAll() {
        // create query
        // "FROM Student" Student is class name mean JPA Entity name not database table name
        // here "FROM Student order by firstName" student class field name not database name
        TypedQuery<Student>theQuery=entityManager.createQuery("FROM Student",Student.class);

        // return query result
        return theQuery.getResultList();
    }

    @Override
    public List<Student> findByLastName(String theLastName) {
        TypedQuery<Student>theQuery=entityManager.createQuery("FROM Student WHERE lastName=:theData",Student.class);

        theQuery.setParameter("theData",theLastName);
        return theQuery.getResultList();
    }
    @Transactional
    @Override
    public void update(Student student) {

        entityManager.merge(student);
    }

    @Override
    @Transactional
    public void delete(Integer studentId) {
        Student student=entityManager.find(Student.class,studentId);
        entityManager.remove(student);
    }

    @Override
    @Transactional
    public int deleteAll() {
        int numRowsDelete=entityManager.createQuery("DELETE FROM Student").executeUpdate();
        return numRowsDelete;
    }
}
