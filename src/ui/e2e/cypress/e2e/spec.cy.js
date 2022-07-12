describe('item list', () => {
  it('shows item list header', () => {
    // arrange
    cy.visit('http://localhost:3001');

    // assert
    cy.get('h2').contains('Items List');
  })

  it('has 10 items', () => {
    // arrange
    cy.visit('http://localhost:3001');

    // assert
    cy.get('#item-list').children().should('have.length', 10)
  })
})