use anchor_lang::prelude::*;

declare_id!("EWLsq62jkAKkj9WQo9gK5U58ZCeu2Rz2M6956vJTmD9");

#[program]
pub mod sol_kit {
    use super::*;

    pub fn initialize(ctx: Context<Initialize>) -> Result<()> {
        msg!("Greetings from: {:?}", ctx.program_id);
        Ok(())
    }
}

#[derive(Accounts)]
pub struct Initialize {}
